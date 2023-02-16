import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public itemIndex: number;
    itemName = '';
    items: Map<any, any> = new Map();

    constructor(private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string) {
        this.items = this.retrieveItemsFromCache();
        this.items.set(id, items);
        this.saveItemToCache(this.items);
    }

    retrieveItemsFromCache() {
        this.items = new Map(JSON.parse(localStorage.getItem(this.itemName)));
        return this.items
    }

    changeMapToArray(map: Map<any, any>) {
        const tableItems: Array<any> = [];
        map.forEach((value) => {
            tableItems.push(this.changeObjectToArray(value));
        });
        return tableItems
    }

    changeObjectToArray(object: Object) {
        const element: Object = {};
        Object.keys(object).map((elementIndex) => {
            typeof object[elementIndex] === 'object' ?
                element[elementIndex] instanceof Date ? element[elementIndex] = new Date(object[elementIndex]) :
                Object.keys(object[elementIndex]).map((elementOfObject) => {
                    element[elementOfObject] = object[elementIndex][elementOfObject];
                }) : element[elementIndex] = object[elementIndex];
        });
        return element;
    }

    clearMultipleItems(selected: any[]) {
        this.items = this.retrieveItemsFromCache();
        this.items.forEach((value, key) => {
            selected.forEach((selectedItem) => {
                if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(selectedItem)) {
                    this.items.delete(key);
                }
            })
        });
        this.saveItemToCache(this.items);
        window.location.reload();
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.itemName, JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    removeItem(id) {
        this.items = this.retrieveItemsFromCache();
        this.items.forEach((value, key) => {
            if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(id)) {
                this.items.delete(key);
            }
        });
        this.saveItemToCache(this.items);
        window.location.reload();
    }

    loadEnumElement(enumElement: any) {
        let tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            tempArray.push(enumElement[index]);
        })
        tempArray = tempArray.filter((item) => item);
        return tempArray;
    }

    editItems(index: number, element: Object) {
        this._router.navigate(['edit' , index], {relativeTo: this._route});
        localStorage.setItem('edit', JSON.stringify(element));
    }

    checkRoute(element: Object) {
        const currentUrl = this._router.url;
        currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase() !== 'new' ? this.isEdit = true : this.isEdit = false;
        this.itemIndex = Number(currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase());
        if (this.isEdit) {
            this.loadElement(this.itemIndex, element);
        }
    }

    loadElement(indexValue: number, element: Object) {
        const editItem = JSON.parse(localStorage.getItem('edit'));
        const elementAsObject = JSON.parse(localStorage.getItem(this.itemName))[indexValue][1];
        if (editItem !== null) {
            Object.keys(elementAsObject).map((elementIndex) => {
                typeof elementAsObject[elementIndex] === 'object' ?
                    elementAsObject[elementIndex] instanceof Date ? element[elementIndex] = new Date(editItem[elementIndex]) :
                            element[elementIndex] = elementAsObject[elementIndex] : element[elementIndex] = editItem[elementIndex];
            });
        }
    }

    saveItemToCache(map: Map<any, any>) {
        localStorage.map = JSON.stringify(Array.from(map.entries()));
        localStorage.setItem(this.itemName, localStorage.map);
    }
}
