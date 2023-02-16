import {Component, Input, OnInit} from '@angular/core';
import {FormService} from '../../services/form.service';

@Component({
    selector: 'app-form',
    templateUrl: './form.component.html',
    styleUrls: ['./form.component.scss'],
    providers: [FormService],
})

export class FormComponent implements OnInit {

    @Input() formElement: any[][];
    public isEdit: boolean;
    public itemIndex: number;

    constructor(public templateService: FormService) { }

    ngOnInit(): void {
        this.templateService.initElement(this.formElement);
        // this.templateService.checkRoute()
    }
}
