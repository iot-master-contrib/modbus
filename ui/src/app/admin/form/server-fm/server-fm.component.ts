import { RequestService } from './../../../request.service';
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormControl, FormGroup, UntypedFormGroup, ValidationErrors, Validators, FormsModule } from '@angular/forms';
import { NzMessageService } from "ng-zorro-antd/message";
@Component({
  selector: 'app-server-fm',
  templateUrl: './server-fm.component.html',
  styleUrls: ['./server-fm.component.scss']
})
export class ServerFmComponent implements OnInit {
  validateForm: UntypedFormGroup;
  deviceList = [];
  constructor(private fb: UntypedFormBuilder, private msg: NzMessageService, private rs: RequestService) {
    this.validateForm = this.fb.group({
      id: [''],
      name: [''],
      desc: [''],
      port: [60000],
      devices: [''],
      period: [60],
      interval: [2],
      protocol: ['rtu']
    });
  }
  ngOnInit(): void {
  }
  show(data: any) {
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
      this.validateForm.patchValue({ port: Number(this.validateForm.value.port) })
      const sendData = Object.assign({}, this.validateForm.value);
      const { id, deviceId } = sendData;
      let url = id ? `server/${id}` : `server/create`
      for (let index = 0; index < this.deviceList.length; index++) {
        const element: { id: string, name: string, product_id: string, slave: number } = this.deviceList[index];
        if (element.id === deviceId) {
          sendData.defaults = [{
            name: element.name,
            product_id: element.product_id,
            slave: element.slave
          }];
          break;
        }
      }
      this.rs.post(url, sendData).subscribe(res => {
        this.msg.success("保存成功")
        this.isVisible = false
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
  reset() {
    this.validateForm.reset();
    for (const key in this.validateForm.controls) {
      if (this.validateForm.controls.hasOwnProperty(key)) {
        this.validateForm.controls[key].markAsPristine();
        this.validateForm.controls[key].updateValueAndValidity();
      }
    }
  }
}
