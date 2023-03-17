import { Component } from '@angular/core';
import {NzContextMenuService,  NzDropdownMenuComponent } from 'ng-zorro-antd/dropdown';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent {
  constructor(private nzContextMenuService: NzContextMenuService){}
  isVisible = false;
v1=false
v2=false
v3=false
v4=false
v5=false

contextMenu($event: MouseEvent, menu: NzDropdownMenuComponent): void {
  this.nzContextMenuService.create($event, menu);
}
selectDropdown(): void {
  // do something
}
  handleOk1(): void { 
    this.v1 = false;
  } 
  handleCancel1(): void { 
    this.v1 = false;
  }
  handleOk2(): void { 
    this.v2 = false;
  } 
  handleCancel2(): void { 
    this.v2 = false;
  }
  handleOk3(): void { 
    this.v3 = false;
  } 
  handleCancel3(): void { 
    this.v3 = false;
  }
  handleOk4(): void { 
    this.v4 = false;
  } 
  handleCancel4(): void { 
    this.v4 = false;
  }
  handleOk5(): void { 
    this.v5 = false;
  } 
  handleCancel5(): void { 
    this.v5 = false;
  }
}
