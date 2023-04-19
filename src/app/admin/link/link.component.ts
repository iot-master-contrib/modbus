import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-link',
  templateUrl: './link.component.html',
  styleUrls: ['./link.component.scss'],
})
export class LinkComponent {
  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();
  }

  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  load() {
    this.loading = true;
    this.rs
      .post('link/search', this.query)
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
    this.rs.get(`link/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.load();
    });
  }

  add() {
    this.router.navigateByUrl(`/admin/create/link`);
  }
  edit(id: number, data: any) {
    const path = `/admin/link/edit/${id}`;
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
    this.router.navigateByUrl('/admin/link/' + id);
  }
  handleToggleStatus(index: number, data: { disabled: boolean, id: number }) {
    const { disabled, id } = data;
    const url = disabled ? `link/${id}/enable` : `link/${id}/disable`;
    this.rs.get(url).subscribe((res) => {
      this.msg.success(`${disabled ? '启用' : '禁用'}成功!`);
      this.load();
    });
  }
}