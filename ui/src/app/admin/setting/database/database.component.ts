import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { RequestService } from '../../../request.service';
import { NzMessageService } from "ng-zorro-antd/message";

@Component({
  selector: 'app-database',
  templateUrl: './database.component.html',
  styleUrls: ['./database.component.scss']
})
export class DatabaseComponent implements OnInit {
  group!: FormGroup;

  constructor(private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService) {this.load()
  }
  switchValue = false;

  ngOnInit(): void {
    // this.rs.get(`config`).subscribe(res => {
    //   //let data = res.data;
    //   this.build(res.data)
    // })

    this.build()
  }

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({
      Type: [obj.type || 'mysql', []],
      URL: ['', [Validators.required]],
      Debug: ['', [Validators.required]],
      LogLevel: ['', [Validators.required]]
    })
  }

  submit() {
    if (this.group.valid) {

      this.rs.post(`config`, this.group.value).subscribe(res => {
        this.msg.success("保存成功")
      })

      return;
    }
    else {
      Object.values(this.group.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });

    }
  }
  loading=false
  query={}
  dbData=[] 
  load(){ 
      this.rs.get(`config/database`).subscribe((res) => {   
       this.dbData=res.data
       this.group.patchValue({LogLevel:res.data.log_level,Type:res.data.type,
        URL:res.data.url  })
      }); 
  }
}
