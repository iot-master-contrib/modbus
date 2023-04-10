import { Component, OnInit, Input, EventEmitter, OnChanges, SimpleChanges } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators } from "@angular/forms";
import { NzMessageService } from "ng-zorro-antd/message";

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
  set listData(data: Array<{ title: string, type?: any, keyName: string }>) {
    const itemObj: any = {};
    for (let index = 0; index < data.length; index++) {
      const { keyName } = data[index];
      itemObj[keyName] = '';
    }
    this.itemObj = itemObj;
    this.constListData = data;
  };
  constructor(
    private msg: NzMessageService,
    private fb: FormBuilder,
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
      properties: this.fb.array(
        obj ? obj.map((prop: any) =>
          this.fb.group(Object.assign(itemObj, prop))
        ) : []
      ),
    })
  }
  handleCopyProperTy(index: number) {
    const item = this.group.get('properties').controls[index];
    this.group.get('properties').controls.splice(index, 0, item);
    this.msg.success("复制成功");
  }
  propertyDel(i: number) {
    this.group.get('properties').controls.splice(i, 1)
  }
  get aliases() {
    return this.group.get('properties') as FormArray;
  }
  add() {
    this.aliases.push(this.fb.group(Object.assign(this.itemObj)));
  }
}