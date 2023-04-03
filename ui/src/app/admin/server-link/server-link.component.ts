import {Component, Input, ViewChild} from '@angular/core';
import {RequestService} from "../../request.service";
import {Router} from "@angular/router";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-server-link',
  templateUrl: './server-link.component.html',
  styleUrls: ['./server-link.component.scss']
})
export class ServerLinkComponent {
  _server = ""

  @Input()
  set server(id: any) {
    this._server = id
    this.query = {server_id: this._server}
    this.load()
  }

  constructor(private router: Router,
              private rs: RequestService,
              private msg: NzMessageService
  ) {
    //this.load();

  }

  @ViewChild('child') child: any
  isVisible!: boolean
  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}

  title!: string
  text!: string

  clientFm(num: number) {
    if (num) this.load()
    this.isVisible = false
  }

  load() {
    this.loading = true
    this.rs.post("link/search", this.query).subscribe(res => {
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    })
  }

  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`link/${id}/delete`).subscribe(res => {
      this.msg.success("删除成功")
      this.isVisible = false;
      this.load()
    })
  }

  add() {
    this.child.reset()
    this.title = "连接端添加"
    this.text = "提交"
    this.isVisible = true
  }

  edit(id: number, data: any) {
    this.title = "连接端修改"
    this.text = "修改"
    this.isVisible = true
    this.child.show(data)
  }

  search(text: any) {
    if (text)
      this.query.filter = {
        server_id: this._server,
        id: text,
      };
    else this.query = {server_id: this._server}
    this.load();
  }

  cancel() {
    this.msg.info('取消删除');
  }

  open(id: string) {
    this.router.navigateByUrl("/admin/link/" + id)
  }
}