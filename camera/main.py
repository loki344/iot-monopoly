import cv2
img = cv2.imread('qrcode.jpg')
det = cv2.QRCodeDetector()
info, box_coordinates, _ = det.detectAndDecode(img)

if box_coordinates is None:
    print('No Code')
else:
    print(info)

if box_coordinates is not None:
    box_coordinates = [box_coordinates[0].astype(int)]
    n = len(box_coordinates[0])
    for i in range(n):
        cv2.line(img, tuple(box_coordinates[0][i]), tuple(box_coordinates[0][(i+1) % n]), (0,255,0), 3)

cv2.imshow('Output', img)
cv2.waitKey(0)
cv2.destroyAllWindows()