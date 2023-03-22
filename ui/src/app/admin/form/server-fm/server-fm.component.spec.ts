import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ServerFmComponent } from './server-fm.component';

describe('ServerFmComponent', () => {
  let component: ServerFmComponent;
  let fixture: ComponentFixture<ServerFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ServerFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ServerFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
