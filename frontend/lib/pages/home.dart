// import 'dart:async';
// import 'dart:io';
// import 'dart:convert';
// import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:frontend/pages/statistics.dart';
import 'package:frontend/pages/version.dart';
import 'package:frontend/pages/manage.dart';

// import 'package:shelf/shelf.dart' as shelf;
// import 'package:shelf/shelf_io.dart' as shelf_io;
//
// // import 'package:shelf_router/shelf_router.dart';
// import 'package:shelf_router/shelf_router.dart';
// import 'package:shelf_router/src/router.dart';
// import 'package:shelf_router/shelf_router.dart';
// import 'package:shelf/shelf.dart';
// import 'package:shelf/shelf_io.dart' as io;

// class Home extends StatelessWidget {
//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       body: Container(
//         color: Colors.black38,
//         height: 500,
//         child: Image.asset('assets/logo.jpg'),
//       ),
//     );
//   }
// }

class Home extends StatefulWidget {
  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {

  // late HttpServer _server;

  List<DataRow> _versionList = [];

  DataRow createRow(String id, String time, String name, String version,
      String status) {
    return DataRow(cells: [
      DataCell(Text(id)),
      DataCell(Text(time)),
      DataCell(Text(name)),
      DataCell(Text(version)),
      DataCell(Text(status))
    ]);
  }

  // @override
  // void initState() {
  //   super.initState();
  //   // startServer();
  // }
  //
  // @override
  // void dispose() {
  //   super.dispose();
  // }

  //
  // Future<void> startServer() async {
  //
  //   var app = Route.get(route) Router;
  //
  //   app.post('/releases', (Request request) {
  //
  //
  //
  //
  //     return Response.ok('received successfully');
  //   });
  //
  //   _server = await io.serve(app, 'localhost', 9090);

    // final router = Router();
    //
    // // Define the JSON endpoint route
    // router.post('/json-endpoint', (shelf.Request request) async {
    //   final requestBody = await request.readAsString();
    //   final jsonBody = jsonDecode(requestBody);
    // }
    //
    //     _router = router;
    //     // Start the HTTP server
    //     final server = await shelf_io.serve(_router.handler, 'localhost', 8080)
  // }

  // void startServer() async {
  //   try {
  //     _server = await HttpServer.bind('localhost', 9090);
  //
  //     await for (var request in _server!) {
  //       if (request.method == 'POST' && request.uri.path == '/releases') {
  //         var content = await request.transform(utf8.decoder as StreamTransformer<Uint8List, dynamic>).join();
  //         var jsonBody = jsonDecode(content);
  //
  //         var id = jsonBody['id'];
  //         var time = jsonBody['time'];
  //         var name = jsonBody['name'];
  //         var version = jsonBody['version'];
  //         var status = jsonBody['status'];
  //
  //         var datarow = createRow(id, time, name, version, status);
  //
  //         setState(() {
  //           _versionList.add(datarow);
  //         });
  //
  //         request.response
  //           ..statusCode = HttpStatus.ok
  //           ..write('JSON object received successfully')
  //           ..close();
  //       } else {
  //         request.response
  //           ..statusCode = HttpStatus.notFound
  //           ..write('Not found')
  //           ..close();
  //       }
  //     }
  //   } catch (e) {
  //     print('Error $e');
  //   }
  // }

  @override
  Widget build(BuildContext context) {
    const double tablePadding = 100;

    return DefaultTabController(
      length: 3,
      child: Scaffold(
        persistentFooterButtons: const [
          Center(
              child: Text('Made With \u{2764} and dedicated to Braam Engall'))
        ],
        appBar: AppBar(
          flexibleSpace: Container(
            decoration: BoxDecoration(
              // LinearGradient
              gradient: LinearGradient(
                // colors for gradient
                colors: [
                  Colors.green,
                  Colors.lightBlue,
                  Colors.blue,
                  Colors.blueAccent,
                  Colors.deepPurple,
                  Colors.purple,
                  Colors.pinkAccent,
                ],
              ),
            ),
          ),
          bottom: const TabBar(
            tabs: [
              Tab(
                icon: Icon(Icons.rocket_launch),
                // text: 'Versions',
                child: Text(
                  'Versions',
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                ),
              ),
              Tab(
                icon: Icon(Icons.text_snippet_outlined),
                // text: 'Manage',
                child: Text(
                  'Manage',
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                ),
              ),
              Tab(
                icon: Icon(Icons.analytics_outlined),
                // text: 'Statistics',
                child: Text(
                  'Statistics',
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                ),
              ),
            ],
          ),
          title: const Center(
              child: Text('Release Version Manager',
                  style: TextStyle(fontSize: 25))),
        ),
        body: Center(
          child: Container(
            padding: const EdgeInsets.fromLTRB(
                tablePadding, tablePadding / 2, tablePadding, 0.0),
            child: TabBarView(
              children: [
                // Version
                VersionPage(versionList: _versionList),
                // Manage
                ManagePage(),
                // Stats
                StatsPage(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}




