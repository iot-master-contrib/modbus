import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PollerFmComponent } from './poller-fm.component';

describe('PollerFmComponent', () => {
  let component: PollerFmComponent;
  let fixture: ComponentFixture<PollerFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PollerFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PollerFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
