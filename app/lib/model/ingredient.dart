import 'package:meta/meta.dart';

class Ingredient {
  final String name;
  final String code;
  final String img;

  const Ingredient({
    @required this.name,
    @required this.code,
    @required this.img,
  });

  factory Ingredient.fromJson(Map<String, dynamic> json) => Ingredient(
        name: json['name'],
        code: json['code'],
        img: json['img'],
  );

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is Ingredient &&
          runtimeType == other.runtimeType &&
          name == other.name &&
          code == other.code &&
          img == other.img;

  @override
  int get hashCode => name.hashCode ^ code.hashCode ^ img.hashCode;
}
