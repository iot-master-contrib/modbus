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
  uploading = false;
  @Input() title = "";
  @Input() downloadHref = "";
  @Input() datum: Array<{ [x: string]: any; name: any; desc: any; created: any; }> = [];
  @Input() uploadObj: any = {};
  @Output() onLoad = new EventEmitter<string>();
  @Output() onSearch = new EventEmitter<string>();
  @Output() onAdd = new EventEmitter<string>();
  constructor(
    private msg: NzMessageService,
    private rs: RequestService
  ) { }
  handleReadCsv(e: any) {
    readCsv(e, this, this.uploadObj.url);
  }
  load() {
    this.onLoad.emit();
  }
  handleChange(info: any): void {
    if (info.file && info.file.response) {
      const res = info.file.response;
      if (!res.error) {
        this.msg.success(`上传成功`);
      } else {
        this.msg.error(`${res.error}`);
      }
    }
    console.log("🚀 ~ file: common-header.component.ts:31 ~ CommonHeaderComponent ~ handleChange ~ info:", info)

  }
}
