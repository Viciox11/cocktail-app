import 'package:cocktail_app/model/drink.dart';
import 'package:cocktail_app/model/ingredient.dart';
import 'package:cocktail_app/page/not_found_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'drink_detail_page.dart';

class DrinkListPage extends StatefulWidget {
  final List<Ingredient> selectedIngredients;

  const DrinkListPage({Key key, @required this.selectedIngredients})
      : super(key: key);

  @override
  _DrinkListState createState() => _DrinkListState();
}

class _DrinkListState extends State<DrinkListPage> {
  var res, drinks;
  List<Drink> drinkToShow = [];

  @override
  void initState() {
    super.initState();
    fetchData();
  }

  fetchData() async {
    res = await rootBundle.loadString('assets/drink_codes.json');
    final drinksTmp = drinksFromJson(res);
    drinksTmp.forEach((drink) {
      var flag = false;
      var i = 0;
      widget.selectedIngredients.forEach((ingredient) {
        if (drink.codes.contains(int.parse(ingredient.code))) {
          flag = true;
          i = i + 1;
        } else {
          flag = false;
        }
      });
      if (flag == true && i == widget.selectedIngredients.length) {
        drinkToShow.add(drink);
      }
    });
    drinkToShow = drinkToShow.toSet().toList();
    drinkToShow.sort((a, b) => a.name.compareTo(b.name));
    if (drinkToShow.isEmpty) {
      Navigator.push(
        context,
        MaterialPageRoute(builder: (context) => NotFoundPage()),
      );
    }
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      constraints: BoxConstraints.expand(),
      decoration: BoxDecoration(
          image: DecorationImage(
              image: AssetImage("assets/images/background.png"),
              fit: BoxFit.cover)),
      child: Scaffold(
        backgroundColor: Colors.transparent,
        appBar: AppBar(
          title: Text("Cocktail App"),
          elevation: 0.0,
          backgroundColor: Colors.transparent,
        ),
        body: Center(
          child: res != null
              ? Column(children: [
                  //ingredientDescription(widget.selectedIngredients),
                  ListView.builder(
                    shrinkWrap: true,
                    itemCount: drinkToShow.length,
                    itemBuilder: (context, index) {
                      var drink = drinkToShow[index];
                      return ListTile(
                        leading: Hero(
                          tag: drink.code,
                          child: CircleAvatar(
                            backgroundImage: NetworkImage(
                              drink.img ??
                                  "http://www.4motiondarlington.org/wp-content/uploads/2013/06/No-image-found.jpg",
                            ),
                          ),
                        ),
                        title: Text(
                          "${drink.name}",
                          style: TextStyle(
                            fontSize: 22,
                            color: Colors.black,
                          ),
                        ),
                        subtitle: Text(
                          "${drink.story}",
                          overflow: TextOverflow.ellipsis,
                          maxLines: 1,
                          style: TextStyle(
                            color: Colors.black,
                          ),
                        ),
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) =>
                                  DrinkDetailPage(drink: drink),
                            ),
                          );
                        },
                      );
                    },
                  )
                ])
              /**/
              : CircularProgressIndicator(backgroundColor: Colors.white),
        ),
      ),
    );
  }
}

  /*Widget ingredientDescription(List<Ingredient> ingredients) {
    return ingredients.length == 1
        ? Row(mainAxisAlignment: MainAxisAlignment.spaceEvenly, children: [
            Hero(
              tag: ingredients.first.code,
              child: CircleAvatar(
                radius: 50.0,
                backgroundImage: AssetImage(
                  "assets/images/1.jpg", //TODO
                ),
              ),
            ),
            Align(
              alignment: Alignment.centerLeft,
              child: Column(children: [
                Text(ingredients.first.name,
                    style: TextStyle(
                      fontSize: 22,
                      color: Colors.black,
                    ),
                    textAlign: TextAlign.left),
                SizedBox(
                  width: 160,
                  child: Text(ingredients.first.descr,
                      overflow: TextOverflow.clip,
                      style: TextStyle(
                        fontSize: 18,
                        color: Colors.black,
                      ),
                      textAlign: TextAlign.left),
                ),
              ]),
            ),
          ])
        : SizedBox(
            height: 1,
          );
  }*/