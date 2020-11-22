# FIR
pkg load signal

f1=0.48;
f2=0.52;
hc=round(fir1(68,f1,'low')*2^15);
printf("Stopband attenuation %d dB\n", 69*22*(f2-f1))
haux=hc';
save filter haux