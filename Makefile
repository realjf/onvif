# ##############################################################################
# # File: Makefile                                                             #
# # Project: onvif                                                             #
# # Created Date: 2023/12/09 11:36:31                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/12/09 11:37:10                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# #                                                                            #
# ##############################################################################


B ?= master
M ?= update

.PHONY: push
push:
	@git add -A && git commit -m ${M} && git push origin ${B}
