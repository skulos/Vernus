import 'package:flutter/material.dart';
import 'package:neon_widgets/neon_widgets.dart';

class Home extends StatelessWidget {

  Widget textWidget(String text) {

    // var characters = text.split('').map((c) => NeonText(
    //   text: c,
    //   spreadColor: Colors.pink,
    //   blurRadius: 20,
    //   textSize: 74,
    //   textColor: Colors.white,
    // ),)
    // .toList();

    // return Row(
    //   children: characters,
    // );
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        NeonText(
          text: 'Ver',
          spreadColor: Colors.pink,
          // blurRadius: 10,
          textSize: 74,
          fontWeight: FontWeight.bold,
          textColor: Colors.white,
        ),
        NeonText(
          text: 'nus',
          spreadColor: Colors.blue,
          // blurRadius: 20,
          textSize: 74,
          fontWeight: FontWeight.bold,
          textColor: Colors.white,
        )
      ],
    );

  }









  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.transparent,
        centerTitle: true,
        title: textWidget('')
      ),
      // body: Container(
      //   color: Colors.black38,
      //   height: 500,
      //   // child: Image.asset('assets/logo.jpg'),
      //   child: Text('Hello World'),
      // ),
    );
  }
}
