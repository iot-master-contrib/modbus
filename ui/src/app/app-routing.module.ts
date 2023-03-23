import { LinkComponent } from './admin/link/link.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminComponent } from './admin/admin.component';
import { SerialComponent } from './admin/serial/serial.component';
import { ServerComponent } from './admin/server/server.component';
import { ClientComponent } from './admin/client/client.component';
import { MapperComponent } from './admin/mapper/mapper.component';
import { PollerComponent } from './admin/poller/poller.component';
import { ProductComponent } from './admin/product/product.component';
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
      { path: 'server', component: ServerComponent },
      { path: 'serial', component: SerialComponent },
      { path: 'client', component: ClientComponent },
      { path: 'mapper', component: MapperComponent },
      { path: 'poller', component: PollerComponent },
      { path: 'link', component: LinkComponent },
      { path: 'product', component: ProductComponent },
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
