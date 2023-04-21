import { RequestService } from '../../../request.service';
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import {
  UntypedFormBuilder,
  UntypedFormControl,
  FormGroup,
  UntypedFormGroup,
  ValidationErrors,
  Validators,
  FormsModule,
  FormBuilder,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { DatePipe } from '@angular/common';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.scss'],
  providers: [DatePipe],
})
export class ProductEditComponent implements OnInit {
  validateForm!: any;
  id: any = 0;
  listData = [{
    title: '名称',
    keyName: 'name'
  }, {
    title: '类型',
    keyName: 'type',
    type: 'select',
    listOfOption: [{
      label: '字',
      value: 'word'
    }, {
      label: '双字',
      value: 'qword'
    }, {
      label: '浮点数',
      value: 'float'
    }, {
      label: '双精度浮点数',
      value: 'double'
    }],
    defaultValue: 'word'
  }, {
    title: '偏移',
    keyName: 'offset',
    type: 'number',
    defaultValue: '0'
  }, {
    title: '大端',
    keyName: 'be',
    type: 'switch',
    defaultValue: true
  }, {
    title: '倍率',
    keyName: 'rate',
    type: 'number',
    defaultValue: 1
  }]
  mode = "new";
  constructor(
    private readonly datePipe: DatePipe,
    private fb: FormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) {
    this.build();
  }

  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.mode = "edit";
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`product/${this.id}`).subscribe((res) => {
        this.build(res.data);
      });
    }
    this.build();
  }

  build(obj?: any) {
    obj = obj || {};
    this.validateForm = this.fb.group({
      id: [obj.id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      mappers: this.fb.array(
        obj.mappers
          ? obj.mappers.map((prop: any) =>
            this.fb.group({
              code: [prop.code || 3, []],
              addr: [prop.addr || 0, []],
              size: [prop.size || 0, []],
              points: this.fb.array(
                prop.points
                  ? prop.points.map((p: any) =>
                    this.fb.group({
                      name: [p.name || '', []],
                      type: [p.type || 'word', []],
                      offset: [p.offset || 0, []],
                      be: [p.be || true, []],
                      rate: [p.rate || 1, []],
                    })
                  )
                  : []
              ),
            })
          )
          : []
      ),
    });
  }


  handleCancel() {
    this.router.navigateByUrl(`/admin/product`);
  }

  submit() {
    this.validateForm.updateValueAndValidity();
    if (this.validateForm.valid) {
      let url = this.id ? `product/${this.id}` : `product/create`;
      if (this.mode === "edit" && !this.validateForm.value.id) {
        this.msg.warning('ID不可为空');
        return;
      }
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/product`);
      });
      return;
    } else {
      Object.values(this.validateForm.controls).forEach((control: any) => {
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

  addMapper() {
    this.validateForm.get('mappers').push(
      this.fb.group({
        code: [3, []],
        addr: [0, []],
        size: [0, []],
        points: this.fb.array([
          this.fb.group({
            name: ['', []],
            type: ['word', []],
            offset: [0, []],
            be: [true, []],
            rate: [1, []],
          }),
        ]),
      })
    );
  }

  drop(mapper: any, event: CdkDragDrop<string[]>): void {
    moveItemInArray(
      mapper.get('points').controls,
      event.previousIndex,
      event.currentIndex
    );
  }

  pointCopy(mapper: any, index: number) {
    const oitem = mapper.get('points').controls[index].value;
    mapper.get('points').insert(index, this.fb.group(oitem));
    this.msg.success('复制成功');
  }
  pointDel(mapper: any, i: number) {
    mapper.get('points').removeAt(i);
  }

  mapperDel(i: number) {
    this.validateForm.get('mappers').removeAt(i);
  }

  pointAdd(mapper: any) {
    mapper.get('points').push(
      this.fb.group({
        name: ['', []],
        type: ['word', []],
        offset: [0, []],
        be: [true, []],
        rate: [1, []],
      })
    );
  }
}
