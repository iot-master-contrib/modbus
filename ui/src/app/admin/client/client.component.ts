import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from "@angular/router";
import { NzMessageService } from "ng-zorro-antd/message";
@Component({
  selector: 'app-client',
  templateUrl: './client.component.html',
  styleUrls: ['./client.component.scss']
})

export class ClientComponent {
  constructor(private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();

  }
  @ViewChild('child') child: any
  isVisible = false
  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  text!: string
  title!: string
  clientFm(num: number) {
    if (num) this.load()
    this.isVisible = false
  }

  load() {
    this.loading = true
    this.rs.post("client/search", this.query).subscribe(res => {
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    })
  }
  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`client/${id}/delete`).subscribe(res => {
      this.msg.success("删除成功")
      this.isVisible = false;
      this.load()
    })
  }
  add() {
    this.child.reset()
    this.title = "客户端添加"
    this.text = "提交"
    this.isVisible = true
  }
  edit(id: number, data: any) {
    this.title = "客户端修改"
    this.text = "修改"
    this.isVisible = true
    this.child.show(data)
  }
  search(text: any) {
    if (text)
      this.query.filter = {
        id: text,
      };
    else this.query = {}
    this.load();
  }
  cancel() { this.msg.info('取消删除'); }
}
