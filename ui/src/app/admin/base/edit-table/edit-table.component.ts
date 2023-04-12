import { Component, OnInit, Input, EventEmitter, OnChanges, SimpleChanges } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators } from "@angular/forms";
import { NzMessageService } from "ng-zorro-antd/message";
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-edit-table',
  templateUrl: './edit-table.component.html',
  styleUrls: ['./edit-table.component.scss']
})

export class EditTableComponent implements OnChanges {
  group!: any;
  itemObj: object = {};
  constListData: any = [];

  @Input() data: any = {};
  @Input()
  set listData(dt: Array<{ title: string, type?: any, keyName: string, defaultValue?: any }>) {
    const itemObj: any = {};
    for (let index = 0; index < dt.length; index++) {
      const { keyName, defaultValue } = dt[index];
      itemObj[keyName] = defaultValue || '';
    }
    this.itemObj = itemObj;
    this.constListData = dt;
  };
  constructor(
    private msg: NzMessageService,
    private fb: FormBuilder
  ) { }
  ngOnChanges(changes: SimpleChanges): void {
    let currentValue = {};
    if (changes['data'] && changes['data'].currentValue) {
      currentValue = changes['data'].currentValue;
    }
    this.build(currentValue)
  }
  build(obj?: any) {
    const itemObj = JSON.parse(JSON.stringify(this.itemObj));
    obj = obj || [];
    this.group = this.fb.group({
      keyName: this.fb.array(
        obj ? obj.map((prop: any) =>
          this.fb.group(Object.assign(itemObj, prop))
        ) : []
      ),
    })
  }
  handleCopyProperTy(index: number) {
    const oitem = this.group.get('keyName').controls[index].value;
    this.aliases.insert(index, this.fb.group(oitem));
    this.msg.success("复制成功");
  }
  propertyDel(i: number) {
    this.group.get('keyName').controls.splice(i, 1)
  }
  get aliases() {
    return this.group.get('keyName') as FormArray;
  }
  add() {
    this.aliases.push(this.fb.group(Object.assign(this.itemObj)));
  }
  drop(event: CdkDragDrop<string[]>): void {
    moveItemInArray(this.group.get('keyName').controls, event.previousIndex, event.currentIndex);
  }
}