import { Component, Input } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { RequestService } from 'src/app/request.service';

@Component({
  selector: 'app-tunnel-edit-retry',
  templateUrl: './tunnel-edit-retry.component.html',
  styleUrls: ['./tunnel-edit-retry.component.scss'],
})
export class TunnelEditRetryComponent {
  constructor(private fb: FormBuilder, private msg: NzMessageService,private rs: RequestService,) {}
  @Input()url!:string
  validateForm!: FormGroup;
  setting() {
    // this.rs.post(this.url, this.validateForm.value).subscribe((res) => {
    //   this.msg.success('设置成功'); 
    // }); 
  }
  patchValue(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      timeout: [mess.timeout || 0],
      enable: [mess.enable || true],
    });
  }
}
