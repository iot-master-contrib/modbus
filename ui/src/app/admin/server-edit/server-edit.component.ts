import { RequestService } from '../../request.service';
import { Component, Input, Output, ViewChild, EventEmitter, OnInit } from '@angular/core';
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
  selector: 'app-server-edit',
  templateUrl: './server-edit.component.html',
  styleUrls: ['./server-edit.component.scss'],
})
export class ServerEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  deviceList = [];
  listData = [{
    title: '从站号',
    keyName: 'slave',
    type: 'number'
  }, {
    title: '名称',
    keyName: 'name'
  }, {
    title: '产品号',
    keyName: 'product_id'
  }]
  defaultEquip: any = [];
  @ViewChild('editTableChild') editTableChild: any;
  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) { }
  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`server/${this.id}`).subscribe((res) => {
        this.patchValue(res.data);
      });
    }
    this.getDeviceList();
    this.patchValue();
  }
  getDeviceList() {
    this.rs.get(`device/list`).subscribe((res) => {
      const { data } = res || {};
      this.deviceList = data || [];
    });
  }
  patchValue(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      id: [mess.id || ''],
      name: [mess.name || ''],
      desc: [mess.desc || ''],
      port: [mess.port || 60000],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || ''],
      defaults: [mess.defaults || []],
    });
    this.defaultEquip = mess.defaults || [];
  }
  handleCancel() {
    this.router.navigateByUrl(`/admin/server`);
  }
  submit() {
    if (this.validateForm.valid) {
      this.validateForm.patchValue({
        port: Number(this.validateForm.value.port),
      });
      const editTableData = this.editTableChild.group.get('keyName').controls.map((item: { value: any; }) => item.value);
      const sendData = Object.assign({}, this.validateForm.value, {
        defaults: editTableData
      });
      let url = this.id ? `server/${this.id}` : `server/create`;
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
