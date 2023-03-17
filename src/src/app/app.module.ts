import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AdminComponent } from './admin/admin.component';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown'; 
import { NzFormModule } from 'ng-zorro-antd/form';  
import { SerialComponent } from './admin/serial/serial.component';
import { ServerComponent } from './admin/server/server.component';
import { ClientComponent } from './admin/client/client.component';
import { MapperComponent } from './admin/mapper/mapper.component';
import { PollerComponent } from './admin/poller/poller.component';
import { NzPageHeaderModule } from 'ng-zorro-antd/page-header';
import { NzSelectModule } from 'ng-zorro-antd/select'; 
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzTreeModule } from 'ng-zorro-antd/tree';
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    AdminComponent, 
    SerialComponent,
    ServerComponent,
    ClientComponent,
    MapperComponent, 
    PollerComponent,  
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    NzLayoutModule,
    NzMenuModule,
    NzIconModule ,
    NzTableModule,
    NzTreeModule,
    NzModalModule,
    NzDividerModule,
    NzButtonModule,
    NzDropDownModule,
    NzFormModule,
    NzPageHeaderModule,
    NzSelectModule 
  ],
  providers: [
    { provide: NZ_I18N, useValue: zh_CN }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
