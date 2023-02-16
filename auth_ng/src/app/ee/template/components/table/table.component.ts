import {Component, Input, OnInit} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {TableDataService} from '../../services/data.service';
import {SelectionModel} from '@angular/cdk/collections';

@Component({
    selector: 'app-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.scss'],
})

export class TableComponent implements OnInit {
    @Input() displayedColumns: any[];
    @Input() isHidden: boolean;
    @Input() selection: SelectionModel<any>

    dataSources: MatTableDataSource<any>;

    constructor(public tableDataService: TableDataService) { }

    ngOnInit() {
        this.dataSources = new MatTableDataSource(this.tableDataService.changeMapToArray(this.tableDataService.retrieveItemsFromCache()));
    }

    allRowsSelected() {
        const totalRowSelected = this.selection.selected.length;
        const totalRow = this.dataSources.data.length;
        return totalRowSelected === totalRow;
    }

    masterToggle() {
        this.allRowsSelected() ?
            this.selection.clear() :
            this.dataSources.data.forEach(element => this.selection.select(element));
    }
}
