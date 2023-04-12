import { Component, Input, Output, EventEmitter } from '@angular/core';
import { readCsv } from "../../../../public";
import { NzMessageService } from 'ng-zorro-antd/message';
import { RequestService } from '../../../request.service';
@Component({
  selector: 'app-common-header',
  templateUrl: './common-header.component.html',
  styleUrls: ['./common-header.component.scss']
})
export class CommonHeaderComponent {
  titleArr: Array<string> = [];
  keyArr: Array<string> = [];
  uploading = false;
  @Input() title = "";
  @Input() datum: Array<{ [x: string]: any; name: any; desc: any; created: any; }> = [];
  @Input() uploadObj: any = {};
  @Input() set listColumns(data: Array<{ title: string; keyName: string }>) {
    for (let index = 0; index < data.length; index++) {
      const item = data[index];
      this.titleArr.push(item.title);
      this.keyArr.push(item.keyName);
    }
  }
  @Output() onLoad = new EventEmitter<string>();
  @Output() onSearch = new EventEmitter<string>();
  @Output() onAdd = new EventEmitter<string>();
  constructor(
    private msg: NzMessageService,
    private rs: RequestService
  ) { }
  handleExport() {
    const listColumns: Array<string> = this.titleArr;
    const data: any[][] = [];
    data.push(listColumns);
    this.datum.forEach((item: { [x: string]: any; name: any; desc: any; created: any; }) => {
      const arr = [];
      for (const key of this.keyArr) {
        if (key === 'created') {
          arr.push(String(item.created));
        } else {
          arr.push(item[key]);
        }
      }
      data.push(arr);
    });
    let csvContent = 'data:text/csv;charset=utf-8,';
    data.forEach(row => { csvContent += row.join(',') + '\n'; });
    let encodedUri = encodeURI(csvContent);
    window.open(encodedUri);
  }
  handleReadCsv(e: any) {
    readCsv(e, this, this.uploadObj.url);
  }
  load() {
    this.onLoad.emit();
  }
}
