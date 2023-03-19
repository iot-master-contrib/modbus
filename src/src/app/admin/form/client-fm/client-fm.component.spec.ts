import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClientFmComponent } from './client-fm.component';

describe('ClientFmComponent', () => {
  let component: ClientFmComponent;
  let fixture: ComponentFixture<ClientFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClientFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClientFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
