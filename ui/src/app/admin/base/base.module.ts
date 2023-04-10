import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from "@angular/forms";

import { NzTableModule } from "ng-zorro-antd/table";
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzFormModule } from "ng-zorro-antd/form";
import { ReactiveFormsModule } from "@angular/forms";
import { DragDropModule } from '@angular/cdk/drag-drop'
import { NzSelectModule } from "ng-zorro-antd/select";
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';

import { SearchBoxComponent } from './search-box/search-box.component';
import { EditTableComponent } from './edit-table/edit-table.component';
@NgModule({
  declarations: [
    SearchBoxComponent,
    EditTableComponent,
  ],
  exports: [
    SearchBoxComponent,
    EditTableComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    NzInputModule,
    NzButtonModule,
    NzTableModule,
    NzFormModule,
    ReactiveFormsModule,
    DragDropModule,
    NzSelectModule,
    NzSpaceModule,
    NzIconModule,
    NzInputNumberModule
  ]
})
export class BaseModule { }
