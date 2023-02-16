import { NgModule } from '@angular/core';

import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatNativeDateModule, MatOptionModule} from '@angular/material/core';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {ReactiveFormsModule} from '@angular/forms';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatDividerModule} from '@angular/material/divider';
import {MatListModule} from '@angular/material/list';
import {MatTableModule} from '@angular/material/table';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatTabsModule} from '@angular/material/tabs';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatMenuModule} from '@angular/material/menu';
import {MatCheckboxModule} from '@angular/material/checkbox';

@NgModule({
    declarations: [],
    imports: [],
    providers: [],
    exports: [
        MatInputModule,
        MatFormFieldModule,
        MatSelectModule,
        MatOptionModule,
        MatIconModule,
        MatButtonModule,
        ReactiveFormsModule,
        MatExpansionModule,
        MatGridListModule,
        MatToolbarModule,
        MatSidenavModule,
        MatDividerModule,
        MatListModule,
        MatTableModule,
        MatPaginatorModule,
        MatTabsModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatMenuModule,
        MatCheckboxModule
    ]
})
export class MaterialModule { }
