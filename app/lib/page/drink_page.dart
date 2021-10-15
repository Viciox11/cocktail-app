import 'package:cocktail_app/model/drink.dart';
import 'package:cocktail_app/widget/search_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'drink_detail_page.dart';

class DrinkPage extends StatefulWidget {
  final List<Drink> drinks;
  final flag;

  const DrinkPage({
    Key key,
    this.drinks = const [],
    this.flag = int,
  }) : super(key: key);

  @override
  _DrinkPageState createState() => _DrinkPageState();
}

class _DrinkPageState extends State<DrinkPage> {
  var res;
  List<Drink> allDrinks;
  String text = '';

  @override
  void initState() {
    super.initState();
    allDrinks = [];
    fetchData();
  }

  bool containsSearchText(Drink drink) {
    final name = drink.name;
    final textLower = text.toLowerCase();
    final drinkLower = name.toLowerCase();

    return drinkLower.contains(textLower);
  }

  fetchData() async {
    res = await rootBundle.loadString('assets/drink_codes.json');
    allDrinks = drinksFromJson(res);
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    fetchData();
    final drinks = allDrinks.where(containsSearchText).toList();
    final style = TextStyle(fontSize: 16);

    return Scaffold(
      appBar: buildAppBar(),
      body: Column(
        children: <Widget>[
          Expanded(
            child: GridView.count(
              crossAxisCount: 2,
              shrinkWrap: true,
              children: drinks.map((drink) {

                return GestureDetector(
                  onTap: () => Navigator.push(
                    context,
                    MaterialPageRoute(builder: (context) => DrinkDetailPage(drink: drink)),
                  ),
                  child: Column(
                    children: <Widget>[
                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceAround,
                        children: [
                          Expanded(
                            child: Column(
                              children: <Widget>[
                                SizedBox(height: 2),
                                Image.network(
                                  drink.img ??
                                      "http://www.4motiondarlington.org/wp-content/uploads/2013/06/No-image-found.jpg",
                                  height: 120,
                                  width: 120,
                                ),
                                ListTile(
                                  onTap: () => Navigator.push(
                                    context,
                                    MaterialPageRoute(
                                        builder: (context) =>
                                            DrinkDetailPage(drink: drink)),
                                  ),
                                  //leading:
                                  title: Text(
                                    drink.name,
                                    style: style,
                                  ),
                                  trailing: Icon(Icons.arrow_forward_outlined),
                                )
                              ],
                            ),
                          )
                        ],
                      ),
                    ],
                  ),
                );}).toList(),
            ),
          ),
        ],
      ),
    );
  }

  Widget buildAppBar() {
    return AppBar(
      title: Text('Seleziona Drink'),
      backgroundColor: Color.fromRGBO(205, 209, 111, 1),
      bottom: PreferredSize(
        preferredSize: Size.fromHeight(60),
        child: SearchWidget(
          text: text,
          onChanged: (text) => setState(() => this.text = text),
          hintText: 'Cerca Drink',
        ),
      ),
    );
  }
}