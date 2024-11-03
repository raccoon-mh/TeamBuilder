import { TabulatorFull } from 'tabulator-tables';

export function tableInit(tabledata, tableid, tablecolumns){

    var table = new TabulatorFull("#"+tableid, {
        height:300,
        layout:"fitColumns",
        data:tabledata,
        columns: tablecolumns,
    });

    table.on("rowClick", function(e, row){ 
        alert("Row " + row.getData().ID + " Clicked!!!!");
    });

}