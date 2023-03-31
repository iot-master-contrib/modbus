import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SearchBoxComponent } from './search-box/search-box.component';
import { FormsModule } from "@angular/forms";
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
@NgModule({
  declarations: [
    SearchBoxComponent,
  ],
  exports: [
    SearchBoxComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    NzInputModule,
    NzButtonModule
  ]
})
export class BaseModule { }
