import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PollerComponent } from './poller.component';

describe('PollerComponent', () => {
  let component: PollerComponent;
  let fixture: ComponentFixture<PollerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PollerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PollerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
