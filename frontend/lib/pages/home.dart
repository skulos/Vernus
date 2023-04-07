import 'package:flutter/material.dart';

class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: Colors.black38,
        height: 500,
        child: Image.asset('assets/logo.jpg'),
      ),
    );
  }
}
