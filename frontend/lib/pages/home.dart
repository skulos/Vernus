import 'package:flutter/material.dart';

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
  List<DataRow> _createRows() {
    var date = DateTime.timestamp().toString();

    return [
      DataRow(cells: [
        DataCell(Text('1')),
        DataCell(Text('$date')),
        DataCell(Text('Person Service')),
        DataCell(Text('v1.2.15')),
        DataCell(Text('Pending'))
      ]),
      DataRow(cells: [
        DataCell(Text('2')),
        DataCell(Text('$date')),
        DataCell(Text('Comment Service')),
        DataCell(Text('v1.2.16')),
        DataCell(Text('Pending'))
      ])
    ];
  }

  DataTable _createDataTable() {
    return DataTable(
      border: TableBorder.all(
        width: 1.0,
      ),
      columns: _createColumns(),
      rows: _createRows(),
      dividerThickness: 5,
      // dataRowHeight: 80,
      showBottomBorder: true,
      headingTextStyle:
          TextStyle(fontWeight: FontWeight.bold, color: Colors.white),
      headingRowColor:
      MaterialStateProperty.resolveWith((states) => Colors.grey),
    );
  }

  List<DataColumn> _createColumns() {
    return [
      // DataColumn(label: Text('ID'), tooltip: 'Book identifier'),
      // DataColumn(label: Text('Book')),
      // DataColumn(label: Text('Author'))
      DataColumn(label: Text('ID')),
      DataColumn(label: Text('Date & Time')),
      DataColumn(label: Text('Name')),
      DataColumn(label: Text('Version')),
      DataColumn(label: Text('Testing Status')),
    ];
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          bottom: const TabBar(
            tabs: [
              Tab(
                icon: Icon(Icons.rocket_launch),
                text: 'Versions',
              ),
              Tab(
                icon: Icon(Icons.text_snippet_outlined),
                text: 'Manage',
              ),
              Tab(
                icon: Icon(Icons.analytics_outlined),
                text: 'Statistics',
              ),
            ],
          ),
          title: const Center(child: Text('Release Version Manager')),
        ),
        body: TabBarView(
          children: [
            // Icon(Icons.rocket_launch, size: 350),
            ListView(
              children: [_createDataTable()],
            ),
            // Manage
            // Stats
          ],
        ),
      ),
    );
  }
}
