import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-device',
  templateUrl: './device.component.html',
  styleUrls: ['./device.component.scss'],
})
export class DeviceComponent {
  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();
  }

  isVisible = false;
  addVisible = false;
  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  listColumns = [{
    title: 'ID',
    keyName: 'id'
  }, {
    title: '产品ID',
    keyName: 'product_id'
  }, {
    title: '名称',
    keyName: 'name'
  }, {
    title: '描述',
    keyName: 'desc'
  }, {
    title: '从站号',
    keyName: 'slave'
  }, {
    title: '创建时间',
    keyName: 'created'
  }];
  uploadObj = {
    url: 'device/create',
    sendKeyNameArr: ['name', 'desc', 'tunnel_id', 'tunnel_id', 'slave']
  }
  load() {
    this.loading = true;
    this.rs
      .post('device/search', this.query)
      .subscribe((res) => {
        this.datum = res.data;
        this.total = res.total;
      })
      .add(() => {
        this.loading = false;
      });
  }
  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`device/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.isVisible = false;
      this.load();
    });
  }
  add() {
    this.router.navigateByUrl(`/admin/create/device`);
  }
  edit(id: number, data: any) {
    const path = `/admin/device/edit/${id}`;
    this.router.navigateByUrl(path);
  }
  search(text: any) {
    if (text)
      this.query.filter = {
        id: text,
      };
    else this.query = {};
    this.load();
  }
  cancel() {
    this.msg.info('取消删除');
  }

  open(id: string) {
    //this.router.navigateByUrl("/admin/device/" + id)
  }
}
