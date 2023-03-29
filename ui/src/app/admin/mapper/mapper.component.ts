import { RequestService } from './../../request.service';
import { Component, ViewChild } from '@angular/core'; 
import {Router} from "@angular/router"; 
import {NzMessageService} from "ng-zorro-antd/message";
@Component({
  selector: 'app-mapper',
  templateUrl: './mapper.component.html',
  styleUrls: ['./mapper.component.scss']
})
export class MapperComponent {
  constructor(private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
) { //this.load();

}
@ViewChild('child') child: any 
input!: string
isVisible!:boolean 
loading = true
datum: any[] = []
total = 1;
pageSize = 20;
pageIndex = 1;
query: any = {}
title!:string
text!:string
  listOfData  = [  
    { 
      id: 53001,
      name: 'John Brown',
      desc:"",
      code:"",
      size:"",
      created:new Date() 
    },
    {
      id: 53002,
      name: 'John Brown',
      desc:"",
      code:"",
      size:"",
      created:new Date() 
    },
    {
      id: 53003,
      name: 'John Brown',
      desc:"",
      code:"",
      size:"",
      created:new Date() 
    }
  ];
  clientFm(num: number) {
    if(num)this.load() 
    this.isVisible = false
  }
  
  load() {
    this.loading = true
    this.rs.post("mapper/search", this.query).subscribe(res=>{
      this.datum = res.data;
      this.total = res.total;
    }).add(()=>{
      this.loading = false;
    })
  }
  delete(index: number, id: number) { 
    this.datum.splice(index, 1);
    this.rs.get(`mapper/${id}/delete`).subscribe(res => {
      this.msg.success("删除成功")
      this.isVisible=false;
      this.load()
    })
  } 
  add(){ 
    this.child.reset()
    this.title="映射表添加"
    this.text="提交"
    this.isVisible=true}
  edit(id: number, data: any) {
    this.title="映射表修改"
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

