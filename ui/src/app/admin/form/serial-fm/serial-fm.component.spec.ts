import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SerialFmComponent } from './serial-fm.component';

describe('SerialFmComponent', () => {
  let component: SerialFmComponent;
  let fixture: ComponentFixture<SerialFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SerialFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SerialFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
