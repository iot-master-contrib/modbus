import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AdminComponent } from './admin/admin.component';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzFormModule } from 'ng-zorro-antd/form';
import { SerialComponent } from './admin/serial/serial.component';
import { ServerComponent } from './admin/server/server.component';
import { ClientComponent } from './admin/client/client.component';
import { DeviceComponent } from './admin/device/device.component';
import { ProductComponent } from './admin/product/product.component';
import { NzPageHeaderModule } from 'ng-zorro-antd/page-header';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzTreeModule } from 'ng-zorro-antd/tree';
import { LinkComponent } from './admin/link/link.component';
import { ClientEditComponent } from './admin/client-edit/client-edit.component';
import { DeviceEditComponent } from './admin/device-edit/device-edit.component';
import { SerialEditComponent } from './admin/serial-edit/serial-edit.component';
import { ServerEditComponent } from './admin/server-edit/server-edit.component';
import { BaseModule } from './admin/base/base.module';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzNotificationModule } from 'ng-zorro-antd/notification';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { LinkEditComponent } from './admin/link-edit/link-edit.component';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { ProductEditComponent } from './admin/product-edit/product-edit.component';
import {NzInputModule} from "ng-zorro-antd/input";
import {NzCollapseModule} from "ng-zorro-antd/collapse";
import {CdkDrag, CdkDropList} from "@angular/cdk/drag-drop";
import {NzSwitchModule} from "ng-zorro-antd/switch";
import {NzCardModule} from "ng-zorro-antd/card";
import {NzAutocompleteModule} from "ng-zorro-antd/auto-complete";
import { TunnelDeviceComponent } from './admin/tunnel-device/tunnel-device.component';
import { ClientDetailComponent } from './admin/client-detail/client-detail.component';
import { ServerDetailComponent } from './admin/server-detail/server-detail.component';
import { LinkDetailComponent } from './admin/link-detail/link-detail.component';
import { SerialDetailComponent } from './admin/serial-detail/serial-detail.component';
import { ServerLinkComponent } from './admin/server-link/server-link.component';
import {NzTabsModule} from "ng-zorro-antd/tabs";
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    AdminComponent,
    SerialComponent,
    ServerComponent,
    ClientComponent,
    DeviceComponent,
    LinkComponent,
    ClientEditComponent,
    DeviceEditComponent,
    SerialEditComponent,
    ServerEditComponent,
    LinkEditComponent,
    ProductComponent,
    ProductEditComponent,
    TunnelDeviceComponent,
    ClientDetailComponent,
    ServerDetailComponent,
    LinkDetailComponent,
    SerialDetailComponent,
    ServerLinkComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    NzLayoutModule,
    NzMenuModule,
    NzIconModule,
    NzTableModule,
    NzTreeModule,
    NzSpaceModule,
    NzMessageModule,
    NzNotificationModule,
    NzModalModule,
    NzDividerModule,
    NzButtonModule,
    NzDropDownModule,
    NzInputNumberModule,
    NzPopconfirmModule,
    NzFormModule,
    NzPageHeaderModule,
    NzSelectModule,
    BaseModule,
    NzInputModule,
    NzCollapseModule,
    CdkDropList,
    CdkDrag,
    NzSwitchModule,
    NzCardModule,
    NzAutocompleteModule,
    NzTabsModule
  ],
  providers: [
    { provide: NZ_I18N, useValue: zh_CN }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
