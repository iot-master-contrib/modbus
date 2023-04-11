import { RequestService } from '../../request.service';
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
  selector: 'app-client-edit',
  templateUrl: './client-edit.component.html',
  styleUrls: ['./client-edit.component.scss'],
})
export class ClientEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
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
      this.rs.get(`client/${this.id}`).subscribe((res) => {
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
      net: [mess.net || 'tcp'],
      addr: [mess.addr || ''],
      port: [mess.port || 1],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || 'rtu'],
      retry_timeout: [mess.retry_timeout || 10],
      retry_maximum: [mess.retry_maximum || 0],
    });
  }

  handleCancel() {
    this.router.navigateByUrl(`/admin/client`);
  }
  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `client/${this.id}` : `client/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/client`);
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
}
