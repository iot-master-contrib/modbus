import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormControl, UntypedFormGroup, ValidationErrors, Validators, FormsModule } from '@angular/forms';
import { RequestService } from './../../../request.service';
import {NzMessageService} from "ng-zorro-antd/message";
@Component({
  selector: 'app-mapper-fm',
  templateUrl: './mapper-fm.component.html',
  styleUrls: ['./mapper-fm.component.scss']
})
export class MapperFmComponent implements OnInit{
  validateForm: UntypedFormGroup;
  constructor(private fb: UntypedFormBuilder,  private msg: NzMessageService,private rs: RequestService) {
    this.validateForm = this.fb.group({
      id: ['' ],
      name: ['' ],
      desc: ['' ],
      code: ['' ], 
      size: ['' ]
    });
  }
  ngOnInit(): void {
    // if(this.data)this.validateForm=this.data
  }
  show(data:any) { 
    this.validateForm.patchValue(data) 
    }  
  @Input() isVisible!: boolean; 
  @Input() title!: string; 
  @Input() text!: string; 
  @Output() back = new EventEmitter()
  handleCancel() {
    this.isVisible = false;  
    this.back.emit(0)
    this.reset();
  }
  handleOk() {
    if (this.validateForm.valid) {
      let id = this.validateForm.value.id;
       let url = id ? `mapper/${id}` : `mapper/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.isVisible = false;
        this.back.emit(1);
      });
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
