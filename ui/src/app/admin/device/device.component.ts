import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from "@angular/router";
import { NzMessageService } from "ng-zorro-antd/message";
@Component({
  selector: 'app-device',
  templateUrl: './device.component.html',
  styleUrls: ['./device.component.scss']
})
export class DeviceComponent {
  constructor(private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) { //this.load();

  }
  @ViewChild('child') child: any
  input!: string
  isVisible = false 
  addVisible=false
  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  clientFm(num: number) {
    this.isVisible = false
    this.addVisible = false 
  }
  load() {
    this.loading = true
    this.rs.post("device/search", this.query).subscribe(res => {
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    })
  }
  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`device/${id}/delete`).subscribe(res => {
      this.msg.success("删除成功")
    })
  }
  read(data: any) {
    this.rs.get(`device/${data.id}/read`).subscribe(res => {
      data.read = true;
    })
  }
  edit(id: number, data: any) {
    this.isVisible = true
    this.child.show(data)
  }
  add(){this.addVisible=true} 
  search() {
    this.query.keyword = {
      name: this.input,
    };
    this.query.skip = 0;
    this.load();
  }
  cancel() { this.msg.info('click cancel'); }
}
