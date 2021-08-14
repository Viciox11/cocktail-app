import 'package:cocktail_app/model/ingredient.dart';
import 'package:cocktail_app/provider/ingredient_provider.dart';
import 'package:cocktail_app/utils.dart';
import 'package:cocktail_app/widget/ingredient_listtile_widget.dart';
import 'package:cocktail_app/widget/search_widget.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'drink_list_page.dart';

class IngredientPage extends StatefulWidget {
  final bool isMultiSelection;
  final List<Ingredient> ingredients;
  final flag;

  const IngredientPage({
    Key key,
    this.isMultiSelection = false,
    this.ingredients = const [],
    this.flag = int,
  }) : super(key: key);

  @override
  _IngredientPageState createState() => _IngredientPageState();
}

class _IngredientPageState extends State<IngredientPage> {
  String text = '';
  List<Ingredient> selectedIngredients = [];

  @override
  void initState() {
    super.initState();

    selectedIngredients = widget.ingredients;
  }

  bool containsSearchText(Ingredient ingredient) {
    final name = ingredient.name;
    final textLower = text.toLowerCase();
    final ingredientLower = name.toLowerCase();

    return ingredientLower.contains(textLower);
  }

  List<Ingredient> getPrioritizedIngredients(List<Ingredient> ingredients) {
    final notSelectedIngredients = List.of(ingredients)
      ..removeWhere((ingredient) => selectedIngredients.contains(ingredient));

    return [
      ...List.of(selectedIngredients)
        ..sort(Utils.ascendingSort),
      ...notSelectedIngredients,
    ];
  }

  @override
  Widget build(BuildContext context) {
    final provider = Provider.of<IngredientProvider>(context);
    final allIngredients = getPrioritizedIngredients(provider.ingredients);
    final ingredients = allIngredients.where(containsSearchText).toList();

    return Scaffold(
      appBar: buildAppBar(),
      body: Column(
        children: <Widget>[
          Expanded(
            child: ListView(
              //crossAxisCount: 2,
              shrinkWrap: true,
              children: ingredients.map((ingredient) {
                final isSelected = selectedIngredients.contains(ingredient);

                return IngredientListTileWidget(
                  ingredient: ingredient,
                  isSelected: isSelected,
                  onSelectedIngredient: selectIngredient,
                );
              }).toList(),
            ),
          ),
          buildSelectButton(context),
        ],
      ),
    );
  }

  Widget buildAppBar() {
    return AppBar(
      title: Text('Seleziona Ingredienti'),
      bottom: PreferredSize(
        preferredSize: Size.fromHeight(60),
        child: SearchWidget(
          text: text,
          onChanged: (text) => setState(() => this.text = text),
          hintText: 'Cerca Ingredienti',
        ),
      ),
    );
  }

  Widget buildSelectButton(BuildContext context) {
    final label = widget.isMultiSelection
        ? 'Seleziona ${selectedIngredients.length} Ingredienti'
        : 'Continue';

    return Container(
      padding: EdgeInsets.symmetric(horizontal: 32, vertical: 12),
      color: Color.fromRGBO(205, 209, 111, 1),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(
          shape: StadiumBorder(),
          minimumSize: Size.fromHeight(40),
          primary: Colors.white,
        ),
        child: Text(
          label,
          style: TextStyle(color: Colors.black, fontSize: 16),
        ),
        onPressed: () {
          if (selectedIngredients.length == 0) {
            final snackBar = SnackBar(
              content: const Text('Seleziona almeno un ingrediente!'),
              action: SnackBarAction(
                label: 'Indietro',
                onPressed: () {
                  // Some code to undo the change.
                },
              ),
            );
            // Find the ScaffoldMessenger in the widget tree
            // and use it to show a SnackBar.
            ScaffoldMessenger.of(context).showSnackBar(snackBar);
          } else if (widget.flag == 1) {
            Navigator.push(
              context, MaterialPageRoute(builder: (context) =>
                DrinkListPage(selectedIngredients: selectedIngredients),
            ),);
          }
        },
      ),
    );
  }

  void selectIngredient(Ingredient ingredient) {
    if (widget.isMultiSelection) {
      final isSelected = selectedIngredients.contains(ingredient);
      setState(() =>
      isSelected
          ? selectedIngredients.remove(ingredient)
          : selectedIngredients.add(ingredient));
    } else {
      Navigator.pop(context, ingredient);
    }
  }
}
