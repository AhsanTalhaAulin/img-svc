

docker exec -i img-db bash <<EOF
mysql -u img_user -p12345678 <<EOF2
    use img_db;
    select* from images;
EOF2

EOF
