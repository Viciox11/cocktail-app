import 'package:cocktail_app/model/drink.dart';
import 'package:flutter/material.dart';

class DrinkDetailPage extends StatelessWidget {
  final Drink drink;

  const DrinkDetailPage({Key key, @required this.drink}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      constraints: BoxConstraints.expand(),
      decoration: BoxDecoration(
          image: DecorationImage(
              image: AssetImage("assets/images/background.png"), fit: BoxFit.cover)),
      child: Scaffold(
        backgroundColor: Colors.transparent,
        appBar: AppBar(
          backgroundColor: Colors.transparent,
          title: Text(drink.name),
          elevation: 0.0,
        ),
        body: Padding(
          padding: EdgeInsets.all(24),
          child: SingleChildScrollView(
            child: Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  Text(
                    "${drink.name}",
                    style: TextStyle(
                      fontSize: 22,
                      color: Colors.black,
                    ),
                  ),
                  SizedBox(
                    height: 20.0,
                  ),
                  Hero(
                    tag: drink.code,
                    child: CircleAvatar(
                      radius: 100.0,
                      backgroundImage: NetworkImage(
                        drink.img,
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 30.0,
                  ),
                  Text(
                    "Storia",
                    style: TextStyle(
                      fontSize: 22,
                      color: Colors.black,
                    ),
                  ),
                  Text("${drink.story}",
                      style: TextStyle(
                        fontSize: 18,
                        color: Colors.black,
                      ),
                      textAlign: TextAlign.center),
                  SizedBox(
                    height: 20.0,
                  ),
                  Text(
                    "Dosi",
                    style: TextStyle(
                      fontSize: 22,
                      color: Colors.black,
                    ),
                  ),
                  Text("${drink.doses}",
                      style: TextStyle(
                        fontSize: 18,
                        color: Colors.black,
                      ),
                      textAlign: TextAlign.center),
                  SizedBox(
                    height: 20.0,
                  ),
                  Text(
                    "Preparazione",
                    style: TextStyle(
                      fontSize: 22,
                      color: Colors.black,
                    ),
                  ),
                  Text("${drink.preparation}",
                      style: TextStyle(
                        fontSize: 18,
                        color: Colors.black,
                      ),
                      textAlign: TextAlign.center),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
