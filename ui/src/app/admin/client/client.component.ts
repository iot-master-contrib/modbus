import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-client',
  templateUrl: './client.component.html',
  styleUrls: ['./client.component.scss'],
})
export class ClientComponent {
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

  load() {
    this.loading = true;
    this.rs
      .post('client/search', this.query)
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
    this.rs.get(`client/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.isVisible = false;
      this.load();
    });
  }
  add() {
    this.router.navigateByUrl(`/admin/create/client`);
  }
  edit(id: number, data: any) {
    const path = `/admin/client/edit/${id}`;
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
    this.router.navigateByUrl('/admin/client/' + id);
  }
}
