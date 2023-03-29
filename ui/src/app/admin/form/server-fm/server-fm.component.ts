import { RequestService } from './../../../request.service';
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormControl, FormGroup, UntypedFormGroup, ValidationErrors, Validators, FormsModule } from '@angular/forms';
 import {NzMessageService} from "ng-zorro-antd/message";
@Component({
  selector: 'app-server-fm',
  templateUrl: './server-fm.component.html',
  styleUrls: ['./server-fm.component.scss']
})
export class ServerFmComponent implements OnInit{
  validateForm: UntypedFormGroup;
  constructor(private fb: UntypedFormBuilder,  private msg: NzMessageService,private rs: RequestService) {
    this.validateForm = this.fb.group({
      id: ['' ],
      name: ['' ],
      desc: ['' ],
      port: [0 ] 
    });
  }
  ngOnInit(): void {
   }
   show(data:any) { 
    this.validateForm.patchValue(data) 
    }    
  @Input() text!: string;  
  @Input() isVisible!: boolean;
  @Input() title!: string; 
  @Output() back = new EventEmitter()
  handleCancel() {
    this.isVisible = false;  
    this.back.emit(0)
    this.reset();
  }
  handleOk() {
 
     if (this.validateForm.valid) { 
      let id=this.validateForm.value.id
      let url =  id ? `server/${ id}` : `server/create`
      this.validateForm.patchValue({port:Number(this.validateForm.value.port)})
      this.rs.post(url, this.validateForm.value).subscribe(res => { 
      this.msg.success("保存成功")
      this.isVisible=false
      this.back.emit(1)
    }) 
      return;
    }
    else {
      Object.values(this.validateForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
  reset(){
    this.validateForm.reset();
    for (const key in this.validateForm.controls) {
      if (this.validateForm.controls.hasOwnProperty(key)) {
        this.validateForm.controls[key].markAsPristine();
        this.validateForm.controls[key].updateValueAndValidity(); 
      }
    }
  }
}
