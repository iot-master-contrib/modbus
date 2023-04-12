import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-common-header',
  templateUrl: './common-header.component.html',
  styleUrls: ['./common-header.component.scss']
})
export class CommonHeaderComponent {
  @Input() title = "";
  @Output() onSearch = new EventEmitter<string>();
  @Output() onAdd = new EventEmitter<string>();
  constructor() { }
}
