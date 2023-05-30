import 'package:flutter/material.dart';

class VersionPage extends StatelessWidget {

  VersionPage({required this.versionList});

  final List<DataRow> versionList;

  DataTable _createDataTable() {
    return DataTable(
      border: TableBorder.all(
        width: 1.0,
      ),
      columns: _createColumns(),
      rows: versionList,
      dividerThickness: 5,
      showBottomBorder: true,
      headingTextStyle:
      const TextStyle(fontWeight: FontWeight.bold, color: Colors.black),
      headingRowColor:
      MaterialStateProperty.resolveWith((states) => Colors.white),
    );
  }

  List<DataColumn> _createColumns() {
    return const [
      DataColumn(label: Text('ID')),
      DataColumn(label: Text('Date & Time')),
      DataColumn(label: Text('Name')),
      DataColumn(label: Text('Version')),
      DataColumn(label: Text('Testing Status')),
    ];
  }

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: [_createDataTable()],
    );
  }
}

