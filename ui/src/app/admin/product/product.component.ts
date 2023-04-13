import { RequestService } from './../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss'],
})
export class ProductComponent {
  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();
  }

  isVisible = false;
  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  uploadObj = {
    url: 'product/create',
    sendKeyNameArr: ['name', 'desc']
  }
  load() {
    this.loading = true;
    this.rs
      .post('product/search', this.query)
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
    this.rs.get(`product/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.isVisible = false;
      this.load();
    });
  }
  add() {
    this.router.navigateByUrl(`/admin/create/product`);
  }
  edit(id: number, data: any) {
    const path = `/admin/product/edit/${id}`;
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
}
