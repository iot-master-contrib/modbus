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
  ) {  this.load();

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
  title!:string
  text!:string
  clientFm(num: number) {
    if(num)this.load() 
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
      this.isVisible=false;
      this.load()
    })
  } 
  add(){ 
    this.child.reset()
    this.title="串口添加"
    this.text="提交"
    this.isVisible=true}
  edit(id: number, data: any) {
    this.title="串口修改"
    this.text="修改"
    this.isVisible = true
    this.child.show(data)
  }
  search() {
    if(this.input)
    this.query.filter = {
      id: this.input,
    };
     else this.query={}
    this.load();
  }
  cancel() { this.msg.info('取消删除'); }
}

