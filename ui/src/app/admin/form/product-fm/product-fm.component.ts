import {RequestService} from '../../../request.service';
import {Component, Input, Output, EventEmitter, OnInit} from '@angular/core';
import {
  UntypedFormBuilder,
  UntypedFormControl,
  FormGroup,
  UntypedFormGroup,
  ValidationErrors,
  Validators,
  FormsModule, FormBuilder,
} from '@angular/forms';
import {NzMessageService} from 'ng-zorro-antd/message';
import {DatePipe} from '@angular/common';
import {CdkDragDrop, moveItemInArray} from "@angular/cdk/drag-drop";

@Component({
  selector: 'app-product-fm',
  templateUrl: './product-fm.component.html',
  styleUrls: ['./product-fm.component.scss'],
  providers: [DatePipe],
})
export class ProductFmComponent implements OnInit {
  validateForm!: any;

  constructor(
    private readonly datePipe: DatePipe,
    private fb: FormBuilder,
    private msg: NzMessageService,
    private rs: RequestService
  ) {
    this.build()
  }


  build(obj?: any) {
    obj = obj || {}
    this.validateForm = this.fb.group({
      id: [obj.id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      mappers: this.fb.array(
        obj.mappers ? obj.mappers.map((prop: any) =>
          this.fb.group({
            code: [prop.code || 3, []],
            addr: [prop.addr || 0, []],
            size: [prop.size || 0, []],
            points: this.fb.array(
              prop.points ? prop.points.map((p: any) =>
                this.fb.group({
                  name: [p.name || '', []],
                  type: [p.type || 'word', []],
                  offset: [p.offset || 0, []],
                  be: [p.be || true, []],
                  rate: [p.rate || 1, []],
                })
              ) : []
            ),
          })
        ) : []
      ),
    })
  }

  ngOnInit(): void {
  }

  show(data: any) {
    this.build(data)
    //this.validateForm.patchValue(data);
  }

  @Input() text!: string;
  @Input() isVisible!: boolean;
  @Input() title!: string;
  @Output() back = new EventEmitter();

  handleCancel() {
    this.isVisible = false;
    this.back.emit(0);
    this.reset();
  }

  handleOk() {
    this.validateForm.updateValueAndValidity()
    if (this.validateForm.valid) {
      let id = this.validateForm.value.id;
      let url = id ? `product/${id}` : `product/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.isVisible = false;
        this.back.emit(1);
      });
      return;
    } else {
      Object.values(this.validateForm.controls).forEach((control: any) => {
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
          })]
        ),
      })
    )
  }

  drop(mapper:any, event: CdkDragDrop<string[]>): void {
    moveItemInArray(mapper.get('points').controls, event.previousIndex, event.currentIndex);
  }

  pointCopy(mapper:any, index: number) {
    const item = mapper.get('points').controls[index];
    mapper.get('points').controls.splice(index, 0, item);
    this.msg.success("复制成功");
  }
  pointDel(mapper:any, i: number) {
    mapper.get('points').removeAt(i)
  }

  mapperDel(i: number) {
    this.validateForm.get("mappers").removeAt(i)
  }

  pointAdd(mapper: any) {
    mapper.get('points').push(this.fb.group({
      name: ['', []],
      type: ['word', []],
      offset: [0, []],
      be: [true, []],
      rate: [1, []],
    }))
  }
}