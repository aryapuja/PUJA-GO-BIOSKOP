PGDMP      &        	        }            bioskop    16.8    16.8     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16398    bioskop    DATABASE     m   CREATE DATABASE bioskop WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en-US';
    DROP DATABASE bioskop;
                postgres    false            �            1259    16400    bioskop    TABLE     �   CREATE TABLE public.bioskop (
    id integer NOT NULL,
    nama text NOT NULL,
    lokasi text NOT NULL,
    rating double precision
);
    DROP TABLE public.bioskop;
       public         heap    postgres    false            �            1259    16399    bioskop_id_seq    SEQUENCE     �   CREATE SEQUENCE public.bioskop_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.bioskop_id_seq;
       public          postgres    false    216            �           0    0    bioskop_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.bioskop_id_seq OWNED BY public.bioskop.id;
          public          postgres    false    215                       2604    16403 
   bioskop id    DEFAULT     h   ALTER TABLE ONLY public.bioskop ALTER COLUMN id SET DEFAULT nextval('public.bioskop_id_seq'::regclass);
 9   ALTER TABLE public.bioskop ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            �          0    16400    bioskop 
   TABLE DATA           ;   COPY public.bioskop (id, nama, lokasi, rating) FROM stdin;
    public          postgres    false    216   b
       �           0    0    bioskop_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.bioskop_id_seq', 3, true);
          public          postgres    false    215                       2606    16407    bioskop bioskop_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.bioskop
    ADD CONSTRAINT bioskop_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.bioskop DROP CONSTRAINT bioskop_pkey;
       public            postgres    false    216            �   W   x�3�t��K-���,V�M,�/��M�I�K�4ѳ�2�)J�+�M,*Q��P���D�5�tvS i�MiI �%�
��y%�Jc���� ��     