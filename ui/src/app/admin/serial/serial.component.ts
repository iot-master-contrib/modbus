import { RequestService } from './../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from "@angular/router";
import { NzMessageService } from "ng-zorro-antd/message";
@Component({
  selector: 'app-serial',
  templateUrl: './serial.component.html',
  styleUrls: ['./serial.component.scss']
})
export class SerialComponent {
  constructor(private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) { //this.load();

  }
  @ViewChild('child') child: any
  input!: string
  isVisible = false
  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  clientFm(num: number) {
    this.isVisible = false
  }
  load() {
    this.loading = true
    this.rs.post("serial/search", this.query).subscribe(res => {
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    })
  }
  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`serial/${id}/delete`).subscribe(res => {
      this.msg.success("删除成功")
    })
  }
  read(data: any) {
    this.rs.get(`serial/${data.id}/read`).subscribe(res => {
      data.read = true;
    })
  }
  edit(id: number, data: any) {
    this.isVisible = true
    this.child.show(data)
  }
  search() {
    this.query.keyword = {
      name: this.input,
    };
    this.query.skip = 0;
    this.load();
  }
  cancel() { this.msg.info('click cancel'); }
}

