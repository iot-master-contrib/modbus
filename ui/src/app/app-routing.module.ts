import { LinkComponent } from './admin/link/link.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminComponent } from './admin/admin.component';
import { SerialComponent } from './admin/serial/serial.component';
import { ServerComponent } from './admin/server/server.component';
import { ClientComponent } from './admin/client/client.component';
import { DeviceComponent } from './admin/device/device.component';
import { ProductComponent } from './admin/product/product.component';
import {ServerDetailComponent} from "./admin/server-detail/server-detail.component";
import {LinkDetailComponent} from "./admin/link-detail/link-detail.component";
import {ClientDetailComponent} from "./admin/client-detail/client-detail.component";
import {SerialDetailComponent} from "./admin/serial-detail/serial-detail.component";
const pages: Routes = [
  {
    path: 'setting',
    loadChildren: () => import('./admin/setting/setting.module').then(m => m.SettingModule)
  }
]
const routes: Routes = [
  {
    path: 'admin',
    component: AdminComponent,
    children: [
      { path: 'product', component: ProductComponent },
      //{ path: 'product/:id', component: ProductDetailComponent },
      { path: 'device', component: DeviceComponent },
      //{ path: 'device/:id', component: DeviceDetailComponent },
      { path: 'server', component: ServerComponent },
      { path: 'server/:id', component: ServerDetailComponent },
      { path: 'link', component: LinkComponent },
      { path: 'link/:id', component: LinkDetailComponent },
      { path: 'serial', component: SerialComponent },
      { path: 'serial/:id', component: SerialDetailComponent },
      { path: 'client', component: ClientComponent },
      { path: 'client/:id', component: ClientDetailComponent },
      ...pages,
      { path: '**', redirectTo: 'server' },
    ]
  },
  {
    path: '**', redirectTo: 'admin'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
