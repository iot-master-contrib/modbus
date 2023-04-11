import {RequestService} from '../../request.service';
import {Component, Input, Output, EventEmitter, OnInit} from '@angular/core';
import {
  UntypedFormBuilder,
  FormGroup,
  ValidationErrors,
  Validators,
  FormsModule,
} from '@angular/forms';
import {NzMessageService} from 'ng-zorro-antd/message';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-serial-edit',
  templateUrl: './serial-edit.component.html',
  styleUrls: ['./serial-edit.component.scss'],
})
export class SerialEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  ports: any = [];

  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.rs.get('serial/ports').subscribe((res) => {
      this.ports = res.data;
    });

    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`serial/${this.id}`).subscribe((res) => {
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
      port_name: [mess.port_name || ''],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || ''],
      retry_timeout: [mess.retry_timeout || 10],
      retry_maximum: [mess.retry_maximum || 0],
    });
  }

  handleCancel() {
    this.router.navigateByUrl(`/admin/serial`);
  }

  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `serial/${this.id}` : `serial/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/serial`);
      });
      return;
    } else {
      Object.values(this.validateForm.controls).forEach((control) => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({onlySelf: true});
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
