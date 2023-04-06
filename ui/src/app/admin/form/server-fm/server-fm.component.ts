import { RequestService } from './../../../request.service';
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import {
  UntypedFormBuilder,
  UntypedFormControl,
  FormGroup,
  UntypedFormGroup,
  ValidationErrors,
  Validators,
  FormsModule,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';
@Component({
  selector: 'app-server-fm',
  templateUrl: './server-fm.component.html',
  styleUrls: ['./server-fm.component.scss'],
})
export class ServerFmComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  deviceList = [];
  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) {}
  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`server/${this.id}`).subscribe((res) => {
        this.patchValue(res.data);
      });
    }
    this.patchValue();
  }
  patchValue(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      id: [mess.id || ''],
      name: [mess.name || ''],
      desc: [mess.desc || ''],
      port: [mess.port || 60000],
      devices: [mess.devices || ''],
      period: [mess.period || 60],
      interval: [mess.interval || 2],
      protocol: [mess.protocol || 'rtu'],
      deviceId: [mess.deviceId || ''],
    });
  }
  handleCancel() {
    this.router.navigateByUrl(`/admin/server`);
  }
  submit() {
    if (this.validateForm.valid) {
      this.validateForm.patchValue({
        port: Number(this.validateForm.value.port),
      });
      const sendData = Object.assign({}, this.validateForm.value);
      const { id, deviceId } = sendData;
      let url = this.id ? `server/${this.id}` : `server/create`;
      for (let index = 0; index < this.deviceList.length; index++) {
        const element: {
          id: string;
          name: string;
          product_id: string;
          slave: number;
        } = this.deviceList[index];
        if (element.id === deviceId) {
          sendData.defaults = [
            {
              name: element.name,
              product_id: element.product_id,
              slave: element.slave,
            },
          ];
          break;
        }
      }
      this.rs.post(url, sendData).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/server`);
      });
      return;
    } else {
      Object.values(this.validateForm.controls).forEach((control) => {
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
