import 'package:cocktail_app/model/ingredient.dart';
import 'package:flutter/material.dart';

class IngredientListTileWidget extends StatelessWidget {
  final Ingredient ingredient;
  final bool isSelected;
  final ValueChanged<Ingredient> onSelectedIngredient;

  const IngredientListTileWidget({
    Key key,
    @required this.ingredient,
    @required this.isSelected,
    @required this.onSelectedIngredient,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final selectedColor = Theme.of(context).primaryColor;
    final style = isSelected
        ? TextStyle(
      fontSize: 18,
      color: selectedColor,
      fontWeight: FontWeight.bold,
    )
        : TextStyle(fontSize: 18);

    return ListTile(
      onTap: () => onSelectedIngredient(ingredient),
      leading: Image.asset(
        'assets/images/1.jpg',
        height: 40,
        width: 40,
        fit: BoxFit.cover
      ),
      title: Text(
        ingredient.name,
        style: style,
      ),
      trailing: isSelected ? Icon(Icons.check, color: selectedColor, size: 26) : null,
    );
  }
}
