import { Component } from '@angular/core';
import { NzContextMenuService, NzDropdownMenuComponent } from 'ng-zorro-antd/dropdown';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent {
  constructor(private nzContextMenuService: NzContextMenuService) { }
  edit!: number
  ary = [false, false
    , false
    , false
    , false]

  // contextMenu($event: MouseEvent, menu: NzDropdownMenuComponent, mes: number): void {
  //   this.edit = mes
  //   this.nzContextMenuService.create($event, menu);
  // }
  // clientFm(num: number) {
  //   this.ary[num] = false
  // }
  selectDropdown(): void {
    this.ary[this.edit] = true
  }

}
