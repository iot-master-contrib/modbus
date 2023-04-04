import { RequestService } from './../../../request.service';
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import {
  UntypedFormBuilder,
  FormGroup,
  ValidationErrors,
  Validators,
  FormsModule,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-serial-fm',
  templateUrl: './serial-fm.component.html',
  styleUrls: ['./serial-fm.component.scss'],
})
export class SerialFmComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  ports: any = [];
  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

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
      port: [mess.port || ''],
      period: [mess.period || 60],
      interval: [mess.interval || 2],
      protocol: [mess.protocol || 'rtu'],
    });
  }

  handleCancel() {
    this.router.navigateByUrl(`/admin/serial`);
  }

  submit() {
    if (this.validateForm.valid) {
      let id = this.validateForm.value.id;
      let url = id ? `serial/${id}` : `serial/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/serial`);
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
