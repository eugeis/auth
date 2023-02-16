import {Injectable} from '@angular/core';
import {Router} from '@angular/router';
import {FormBuilder, FormControl, Validators} from '@angular/forms';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class FormService {

    protected elementNameWithValue: any[][];

    form = this.fb.group({});

    formArrayValue = this.fb.array([])

    formArrayName = [];

    formArrayType = [];

    formEnumElement = [];

    index = 0;

    dateNow = this.formatDate(new Date());

    selected = [];

    public isEdit: boolean;
    public itemIndex: number;

    constructor(public fb: FormBuilder, private tableDataService: TableDataService, private _router: Router) {
    }

    formatDate(date: Date) {
        return [date.getFullYear(), '0' + (date.getMonth() + 1), date.getDate()].join('-');
    }

    changeValue(input: string, elementName: string) {
        this.form.get(elementName).setValue(input);
    }

    init() {
        this.initElement(this.elementNameWithValue);
    }

    initElement(elements: any[][]) {
        elements.forEach(el => {
            const elementName = el.at(0);
            const elementType = el.at(1);
            if (elementType === 'string') {
                this.form.addControl(elementName, new FormControl('', [Validators.required]));
                this.formEnumElement.push([]);
            } else if (elementType === 'number') {
                this.form.addControl(elementName, new FormControl(0, [Validators.required]));
                this.formEnumElement.push([]);
            } else if (elementType === 'boolean') {
                this.form.addControl(elementName, new FormControl(false, [Validators.required]));
                this.formEnumElement.push([]);
            } else if (elementType === 'datetime') {
                this.form.addControl(elementName, new FormControl(this.dateNow, [Validators.required]));
                this.formEnumElement.push([]);
            } else if (elementType === 'enum') {
                const enumElement = el.at(2);
                let tempArray = [];
                Object.keys(enumElement).map((element, index) => {
                    tempArray.push(enumElement[index]);
                })
                tempArray = tempArray.filter((item) => item);
                this.formEnumElement.push(tempArray);
                this.form.addControl(elementName, new FormControl(0, [Validators.required]));
            } else {
                this.form.addControl(elementName, new FormControl('', [Validators.required]));
                this.formEnumElement.push([]);
            }
            this.formArrayName.push(elementName);
            this.formArrayType.push(elementType);
        });
    }
}
